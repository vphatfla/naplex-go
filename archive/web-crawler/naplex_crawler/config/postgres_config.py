import os
import logging
import psycopg2
from psycopg2 import pool
from psycopg2.extras import DictCursor
from typing import Optional, List, Tuple, Any, Dict, Union
from contextlib import contextmanager
from dotenv import load_dotenv

# Configure logging
logging.basicConfig(
    level=logging.INFO,
    format='%(asctime)s - %(name)s - %(levelname)s - %(message)s',
    handlers=[
        logging.FileHandler("database.log"),
        logging.StreamHandler()
    ]
)

class PostgresDBConnection:
    """
    A class to manage persistent connections to a PostgreSQL database.
    Uses connection pooling for efficiency and environment variables for security.
    """
    _instance = None
    _connection_pool = None
    _logger = logging.getLogger('PostgresDBConnection')

    def __new__(cls, *args, **kwargs):
        """
        Implements the singleton pattern to ensure only one instance of the connection pool exists.
        """
        if cls._instance is None:
            cls._instance = super(PostgresDBConnection, cls).__new__(cls)
            cls._instance._initialized = False
        return cls._instance

    def __init__(self, min_connections: int = 1, max_connections: int = 10):
        """
        Initializes the database connection pool if it hasn't been initialized already.
        """
        if self._initialized:
            return

        self._logger.info("Initializing PostgreSQL connection pool")
        # Load environment variables
        load_dotenv()

        # Get database configuration from environment variables
        self.db_config = {
            'host': os.getenv('DB_HOST', 'localhost'),
            'port': os.getenv('DB_PORT', '5432'),
            'database': os.getenv('DB_NAME'),
            'user': os.getenv('DB_USER'),
            'password': os.getenv('DB_PASSWORD')
        }

        # Validate configuration
        missing_vars = [key for key, value in self.db_config.items() if value is None]
        if missing_vars:
            self._logger.error(f"Missing environment variables: {', '.join(missing_vars)}")
            raise ValueError(f"Missing required environment variables: {', '.join(missing_vars)}")

        # Create connection pool
        try:
            self._connection_pool = pool.ThreadedConnectionPool(
                min_connections,
                max_connections,
                **self.db_config
            )
            self._logger.info(f"Connection pool created with {min_connections}-{max_connections} connections")
            self._initialized = True
        except Exception as e:
            self._logger.error(f"Failed to create connection pool: {str(e)}")
            raise

    def get_connection(self) -> tuple:
        """
            Returns:
            tuple: (connection, cursor)
        """
        if not self._connection_pool:
            self._logger.error("Connection pool is not initialized")
            raise RuntimeError("Database connection pool is not initialized")

        try:
            connection = self._connection_pool.getconn()

            cursor = connection.cursor()
            cursor.execute("SELECT 1")
            cursor.fetchone()
            cursor.close()

            cursor = connection.cursor()
            return connection, cursor
        except Exception as e:
            self._logger.error(f"Error getting connection from pool: {str(e)}")
            raise

    def release_connection(self, connection):
        """
        Returns a connection back to the pool.

        Args:
            connection: The connection to return to the pool
        """
        if self._connection_pool:
            self._connection_pool.putconn(connection)

    @contextmanager
    def connection(self):
        """
        Context manager for database connections.

        Example:
            with db.connection() as conn:
                with conn.cursor() as cursor:
                    cursor.execute("SELECT * FROM users")
        """
        connection = None
        try:
            connection, _ = self.get_connection()
            yield connection
        finally:
            if connection:
                self.release_connection(connection)

    def execute_query(self, query: str, params: Optional[tuple] = None,
                     fetch: bool = True, as_dict: bool = False) -> Union[List[Tuple], List[Dict], int]:
        """
        Executes a SQL query and returns the results.

        Args:
            query: The SQL query to execute
            params: Optional parameters for the query
            fetch: Whether to fetch and return results (for SELECT queries)
            as_dict: If True, returns results as dictionaries instead of tuples

        Returns:
            For SELECT queries (fetch=True): List of result tuples or dicts
            For other queries (fetch=False): Number of affected rows
        """
        connection = None
        cursor = None
        try:
            connection, cursor = self.get_connection()

            if as_dict:
                cursor.close()
                cursor = connection.cursor(cursor_factory=DictCursor)

            self._logger.debug(f"Executing query: {query}")
            cursor.execute(query, params)

            if fetch and cursor.description is not None:  # It's a SELECT query
                results = cursor.fetchall()
                self._logger.debug(f"Query returned {len(results)} rows")
                return results
            else:  # It's an INSERT, UPDATE, DELETE query
                affected_rows = cursor.rowcount
                connection.commit()
                self._logger.debug(f"Query affected {affected_rows} rows")
                return affected_rows

        except Exception as e:
            if connection:
                connection.rollback()
            self._logger.error(f"Error executing query: {str(e)}")
            raise
        finally:
            if cursor:
                cursor.close()
            if connection:
                self.release_connection(connection)

    def execute_transaction(self, queries: List[Dict[str, Any]]) -> List[Any]:
        """
        Executes multiple queries in a single transaction.

        Args:
            queries: List of dictionaries with keys:
                     - 'query': SQL query string
                     - 'params': Query parameters (optional)
                     - 'fetch': Whether to fetch results (optional, default True)
                     - 'as_dict': Return results as dictionaries (optional, default False)

        Returns:
            List of results for each query in the same order
        """
        connection = None
        cursor = None
        results = []

        try:
            connection, cursor = self.get_connection()

            for query_dict in queries:
                query = query_dict['query']
                params = query_dict.get('params')
                fetch = query_dict.get('fetch', True)
                as_dict = query_dict.get('as_dict', False)

                if as_dict:
                    cursor.close()
                    cursor = connection.cursor(cursor_factory=DictCursor)

                self._logger.debug(f"Executing transaction query: {query}")
                cursor.execute(query, params)

                if fetch and cursor.description is not None:
                    query_result = cursor.fetchall()
                    results.append(query_result)
                else:
                    results.append(cursor.rowcount)

            connection.commit()
            self._logger.info(f"Transaction completed successfully with {len(queries)} queries")
            return results

        except Exception as e:
            if connection:
                connection.rollback()
            self._logger.error(f"Transaction failed: {str(e)}")
            raise
        finally:
            if cursor:
                cursor.close()
            if connection:
                self.release_connection(connection)

    def close_all_connections(self):
        """
        Closes all connections in the pool and shuts down the pool.
        Should be called when the application is shutting down.
        """
        if self._connection_pool:
            self._connection_pool.closeall()
            self._logger.info("All database connections closed")
            self._connection_pool = None
            self._initialized = False
