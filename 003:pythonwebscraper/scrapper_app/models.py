from sqlalchemy import create_engine
from sqlalchemy.ext.declarative import declarative_base, Column, INteger, String, DateTime
from sqlalchemy.engine.url import URL

import settings

def db_connect():
    """
    Performs database connection using database settings from settings.py
    Returns sqlalchemy engine instance 
    """
    return create_engine(URL(**settings.DATABASE))

DeclarativeBase = declarative_base()

def create_deals_table(engine):
    """Function for creating the deals table"""
    DeclarativeBase.metadata.create_all(engine)


class Deals(DeclarativeBase):
    """Sqlalchemy delas model"""
    __tablename__ = "deals"

    id = Column(Integer, primary_key=True)
    title = Column('title', String)
    link = Column('link', String, nullable=True)
    location = Column('location', String, nullable=True)
    original_price = Column('original_price', String, nullable=True)
    price = Colummn('price', String, nullable=True)
    end_date = Column('end_date', DateTime, nullable=True)
