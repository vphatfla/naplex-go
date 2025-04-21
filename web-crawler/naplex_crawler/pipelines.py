# Define your item pipelines here
#
# Don't forget to add your pipeline to the ITEM_PIPELINES setting
# See: https://docs.scrapy.org/en/latest/topics/item-pipeline.html


# useful for handling different item types with a single interface
import json
import os
from datetime import datetime
from scrapy import signals
from naplex_crawler.config import PostgresDBConnection

class JsonWritePipeline:
    """Pipeline to store questions as JSON"""

    def open_spider(self, spider):
        """
        Initialize postgesdb connection pool
        """
        self.logger = spider.logger
        self.logger.info(f'open_spider, init db connection pool')
        self.db = PostgresDBConnection(min_connections=2, max_connections=20)

        """
        self.logger = spider.logger
        self.logger.info('Creating data directory')

        # Get the directory of the current file (pipelines.py)
        current_dir = os.path.dirname(os.path.abspath(__file__))

        # Navigate up two directories to get to naplex-go directory
        # From: ~/workplace/naplex-go/web-crawler/naplex_crawler
        # To:   ~/workplace/naplex-go
        naplex_go_dir = os.path.abspath(os.path.join(current_dir, '..', '..'))

        data_dir = os.path.join(naplex_go_dir, 'data')
        self.logger.info(f'Data directory path: {data_dir}')

        os.makedirs(data_dir, exist_ok=True)

        timestamp = datetime.now().strftime('%Y%m%d_%H%M%S')
        self.json_file = open(f'{data_dir}/naplex_questions_{timestamp}.json', 'w')
        self.text_file = open(f'{data_dir}/naplex_questions_{timestamp}.txt', 'w')
        self.items = []
        """

    def close_spider(self, spider):
        """ Save questions when spider is closed
        self.logger.info(f'close spider w len item {len(self.items)}')
        json.dump(self.items, self.json_file, indent=2, ensure_ascii=False)
        self.json_file.close()
        for i, item in enumerate(self.items, 1):
            self.text_file.write(f"Question {i}\n")
            self.text_file.write('=' * 80 + '\n')  # Header separator

            # Write each field in a key: value format
            for key, value in item.items():
                # Handle multiline values (like question content)
                if isinstance(value, str) and '\n' in value:
                    self.text_file.write(f"{key}:\n{value}\n\n")
                else:
                    self.text_file.write(f"{key}: {value}\n")

            self.text_file.write('\n' + '-' * 80 + '\n\n')
        """
        self.logger.info(f'Closing spider, closing connection')
        self.db.close_all_connections()

    def process_item(self, item, spider):
        """Process logic"""
        query = """
        INSERT INTO naplex_data.raw_questions (title, raw_question, link)
        VALUES (%s, %s, %s)
        """

        # Execute the query with parameters from the item dictionary
        # Map the raw_text field from item to raw_question column in DB
        params = (item.get('title'), item.get('raw_text'), item.get('link'))

        try:
            # Execute the query and return the affected rows count
            # Set fetch=False since this is an INSERT statement
            self.logger.info(f'Writing to db item :: {item.get('title')} at {item.get('link')}')
            result = self.db.execute_query(query, params, fetch=False)
            self.logger.info(f'Result writing {result}')
        except Exception as e:
            # Log the error and re-raise
            self.logger.info(f"Error inserting item into raw_questions: {str(e)}")
            raise

