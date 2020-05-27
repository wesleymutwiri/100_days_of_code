BOT_NAME = 'livingsocial'
SPIDER_MODULES = ['scrapper_app.spiders']
ITEM_PIPELINES = ['scrapper_app.pipelines.LivingSocialPipeline']

DATABASE = {
    'drivername': 'postgres',
    'host': 'localhost',
    'port': '5432',
    'username': 'wesley',
    'password': '',
    "database": 'scrape'
}

