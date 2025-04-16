import logging
from dotenv import load_dotenv
import os
from google import genai

logging.basicConfig(
        level=logging.INFO,  # Set the logging level
        format='%(asctime)s>>>    %(message)s',
        handlers=[
            logging.StreamHandler(),  # This adds console output
            ]
        )
logger = logging.getLogger(__name__)

load_dotenv()


def request_gemini(client) -> str:
    res = client.models.generate_content(
            model="gemini-2.0-flash", contents="Return a dummy test data in a json format with 3 fields"
            )
    return res

def main():
    logger.info('AI Cleaner begining')
    GOOGLE_API_KEY = os.environ.get('GOOGLE_API_KEY')

    client = genai.Client(api_key=GOOGLE_API_KEY)
    res = request_gemini(client)
    print(res.text)
if __name__=="__main__":
    main()
