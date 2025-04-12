# Define your item pipelines here
#
# Don't forget to add your pipeline to the ITEM_PIPELINES setting
# See: https://docs.scrapy.org/en/latest/topics/item-pipeline.html


# useful for handling different item types with a single interface
import json
import os
from datetime import datetime
from scrapy import signals

# useful for handling different item types with a single interface
from itemadapter import is_item, ItemAdapter


class JsonWritePipeline:
    """Pipeline to store questions as JSON"""

    def open_spider(self, spider):
        self.logger = spider.logger
        self.logger.info(f'Making dir ')
        current_dir = os.path.dirname(os.path.abspath(__file__))
        output_dir = os.path.join(current_dir, 'output')
        self.logger.info(f'output dir = {output_dir}')
        os.makedirs(output_dir, exist_ok=True)

        timestamp = datetime.now().strftime('%Y%m%d_%H%M%S')
        self.json_file = open(f'{output_dir}/naplex_questions_{timestamp}.json', 'w')
        self.text_file = open(f'{output_dir}/naplex_questions_{timestamp}.txt', 'w')
        self.items = []

    def close_spider(self, spider):
        """ Save questions when spider is closed"""
        self.logger.info(f'close spider w len item {len(self.items)}')
        json.dump(self.items, self.json_file, indent=2, ensure_ascii=False)
        self.json_file.close()

    def process_item(self, item, spider):
        """Process logic"""
        self.items.append(dict(item))
        return item

