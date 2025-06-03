import scrapy
from urllib.parse import urljoin
from scrapy.loader import ItemLoader
from naplex_crawler.items import QuestionItem
import re
import brotli
from w3lib.html import remove_tags


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
        self.logger.info(f"Doing request with url = {url} \n cookies  = {cookies}")
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
            yield scrapy.Request(url=urljoin(response.url,next_page), callback=self.parse_page, headers=self.headers, cookies=cookies)
        # self.start_requests(url=next_page, cookies=cookies)
    def parse_question(self, response):
        self.logger.info("Start parsing question")
        """Parse func for individual question"""
        loader = ItemLoader(item=QuestionItem(), response=response)

        title = response.css('h1.content__headline::text').get().split(':',1)[1].strip()
        self.logger.info(f'Title = {title}')
        loader.add_value('title', title)

        loader.add_value('link', response.url)

        article_div = response.css('div[data-zapnito-article]')

        if not article_div:
            self.logger.warning("No div with data-zapnito-article attribute found")
            return

        # Get all paragraph texts, handling nested tags
        clean_paragraphs = []

        for p in article_div.css('p'):
            # Get the HTML content of the paragraph with all nested elements
            p_html = p.get()

            if p_html:
                # Remove all HTML tags while preserving their text content
                clean_text = remove_tags(p_html)

                # Normalize whitespace (remove extra spaces, tabs, newlines)
                clean_text = re.sub(r'\s+', ' ', clean_text).strip()

                if clean_text:  # Only include non-empty paragraphs
                    clean_paragraphs.append(clean_text)

        loader.add_value('raw_text', '\n'.join(clean_paragraphs))

        return loader.load_item()
