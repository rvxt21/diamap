# This is an auto-generated Django model module.
# You'll have to do the following manually to clean this up:
#   * Rearrange models' order
#   * Make sure each model has one field with primary_key=True
#   * Make sure each ForeignKey and OneToOneField has `on_delete` set to the desired behavior
#   * Remove `managed = False` lines if you wish to allow Django to create, modify, and delete the table
# Feel free to rename the models, but don't rename db_table values or field names.
import os
from django.db import models

from django.conf import settings


class Continent(models.Model):
    continentid = models.AutoField(primary_key=True)
    continentname = models.CharField(max_length=20, blank=True, null=True)

    def __str__(self):
        return self.continentname
    
    class Meta:
        managed = True
        db_table = 'continents'


class Countrie(models.Model):
    countryid = models.AutoField(primary_key=True)
    countryname = models.CharField(max_length=100)
    continentid = models.ForeignKey('Continent', models.DO_NOTHING, db_column='continentid', blank=True, null=True)
    flag_emoji = models.CharField(max_length=10, blank=True, null=True)
    
    def __str__(self):
        return self.countryname
    
    class Meta:
        managed = True
        db_table = 'countries'


class Podcast(models.Model):
    podcastid = models.AutoField(primary_key=True)
    podcastname = models.CharField(max_length=100)
    duration = models.CharField(max_length=20)
    podcastdescription = models.TextField()
    audioname = models.CharField(max_length=100)
    photoroute = models.ImageField(upload_to='photo/', blank=True, null=True) 
    audioroute = models.FileField(upload_to='audio/', blank=True, null=True) 
    countryid = models.ForeignKey('Countrie', models.DO_NOTHING, db_column='countryid')

    class Meta:
        managed = True
        db_table = 'podcasts'
