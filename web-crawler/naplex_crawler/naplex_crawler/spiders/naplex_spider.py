import scrapy
from scrapy.loader import ItemLoader
from naplex_crawler.items import QuestionItem
import re


class NaplexSpider(scrapy.Spider):
    """Spider for scraping questions"""
    name = 'naplex'
    allow_domains = ['accessmedicinenetwork.com']
    start_urls = ['https://www.accessmedicinenetwork.com/channels/naplex-review-question-of-the-week']

    def start_requests(self):
        headers = {'Accept': 'text/html,application/xhtml+xml,application/xml',
                   'User-Agent': 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/18.4 Safari/605.1.15'}
        for url in self.start_urls:
            yield scrapy.Request(url, headers=headers)

    def parse(self, response):
        """Parse the list of follow page w/ question links"""
        for article in response.css('article.content-card'):
            headline = article.css('h1.content-card_headline a')
            if headline:
                question_url = headline.attrib('href')
                yield response.follow(question_url, self.parse_question)

        next_page = response.css('a.next_page::attr(href)').get()
        if next_page:
            yield response.follow(next_page, self.parse)

    def parse_question(self, response):
        """Parse func for individual question"""
        loader = ItemLoader(item=QuestionItem(), response=response)

        loader.add_css('title', 'h1.content_headeline::text')

        content_div = response.css('div[data-zapnito-article].content_body')
        if content_div:
            paragraphs = []
            for p in content_div.css('p'):
                span = p.css('span')
                if span:
                    text = span('::text').getall()
                    if text:
                        clean_text = ' '.join([t.strip() for t in text if t.strip()])
                        clean_text = re.sub(r'\s+', ' ', clean_text).strip()
                        if clean_text:
                            paragraphs.append(clean_text)

            content = '\n\n'.join(paragraphs)
            loader.load_value('raw_text', content)

        return loader.load_item()
