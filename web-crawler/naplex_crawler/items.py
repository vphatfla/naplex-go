# Define here the models for your scraped items
#
# See documentation in:
# https://docs.scrapy.org/en/latest/topics/items.html

import scrapy

class QuestionItem(scrapy.Item):
    """ Struct of individual naplex item"""
    title = scrapy.Field()
    question = scrapy.Field()
    list_choice = scrapy.Field()
    correct_answer = scrapy.Field()
    raw_text = scrapy.Field()
    link = scrapy.Field()
class NaplexCrawlerItem(scrapy.Item):
    # define the fields for your item here like:
    # name = scrapy.Field()
    pass
