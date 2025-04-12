import scrapy
from urllib.parse import urljoin
from scrapy.loader import ItemLoader
from naplex_crawler.items import QuestionItem
import re
import brotli

class NaplexSpider(scrapy.Spider):
    """Spider for scraping questions"""
    name = 'naplex-spider'
    allow_domains = ['accessmedicinenetwork.com']
    headers = {
                'User-Agent': 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36',
                'Accept': 'text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,*/*;q=0.8',
                'Accept-Language': 'en-US,en;q=0.5',
                'Accept-Encoding': 'gzip, deflate, br',
                'Connection': 'keep-alive',
                'Upgrade-Insecure-Requests': '1',
                'Sec-Fetch-Dest': 'document',
                'Sec-Fetch-Mode': 'navigate',
                'Sec-Fetch-Site': 'none',
                'Sec-Fetch-User': '?1',
                'Cache-Control': 'max-age=0',
                }
    start_url = 'https://www.accessmedicinenetwork.com/channels/naplex-review-question-of-the-week'

    def start_requests(self, url='', cookies={}):
        if url == '':
            url = self.start_url
        self.logger.info(f"Doing request with url = {url}cookies  = {cookies}")
        yield scrapy.Request(url=url, callback=self.parse_page, headers=self.headers, cookies=cookies)

    def parse_page(self, response):
        """Start parsing the page in url, but first must check for encoding"""
        self.logger.info("Start parse page")
        content_encoding = response.headers.get('Content-Encoding')
        result_response = response
        # Decompress brotli if needed
        if content_encoding:
            try:
                print(response.body)
                decompressed_body = brotli.decompress(response.body)
                result_response = response.replace(
                    body=decompressed_body,
                    encoding='utf-8'  # Set the encoding as per the Content-Type header
                )
            except Exception as e:
                self.logger.error(f"Error decompressing Brotli content: {e}")
                return

        cookies_array = result_response.headers.getlist('Set-Cookie')[0].decode("utf-8").split(";")[0].split("=")
        cookies = {}
        self.logger.info(f"Cookies array = {cookies_array}")
        if len(cookies_array) >= 2:
            # Create a dictionary with the first element as key and second as value
            cookies = {cookies_array[0]: cookies_array[1]}
        # Parse the questions
        for article in result_response.css('article.content-card'):
            headline = article.css('h1.content-card__headline')
            if headline:
                question_url = headline.css('a::attr(href)').get()
                self.logger.info(f'Requesting question {question_url}')
                yield scrapy.Request(urljoin(response.url, question_url),
                                             callback=self.parse_question,
                                             headers=self.headers,
                                             cookies=cookies)

        next_page = response.css('a.next_page::attr(href)').get()
        if next_page:
            self.logger.info(f'Requesting next page {next_page}')
            return
        #yield scrapy.Request(url=urljoin(response.url,next_page), callback=self.parse_page, headers=self.headers, cookies=cookies)
        # self.start_requests(url=next_page, cookies=cookies)
    def parse_question(self, response):
        self.logger.info("Start parsing question")
        """Parse func for individual question"""
        loader = ItemLoader(item=QuestionItem(), response=response)

        loader.add_css('title', 'h1.content_headeline::text')

        content_div = response.css('div[data-zapnito-article]')
        self.logger.info(f'Content_div {content_div}')
        if content_div:
            paragraphs = []
            for p in content_div.css('p'):
                self.logger.info(f'p = {p}')
                for span in p.css('span'):
                    text = span.css('::text').get()
                    self.logger.info(f'text = {text}')
                    if text:
                        clean_text = ' '.join([t.strip() for t in text if t.strip()])
                        clean_text = re.sub(r'\s+', ' ', clean_text).strip()
                        if clean_text:
                            paragraphs.append(clean_text)

            content = '\n\n'.join(paragraphs)
            loader.add_value('raw_text', content)

        return loader.load_item()
