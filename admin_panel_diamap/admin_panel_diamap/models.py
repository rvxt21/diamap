# This is an auto-generated Django model module.
# You'll have to do the following manually to clean this up:
#   * Rearrange models' order
#   * Make sure each model has one field with primary_key=True
#   * Make sure each ForeignKey and OneToOneField has `on_delete` set to the desired behavior
#   * Remove `managed = False` lines if you wish to allow Django to create, modify, and delete the table
# Feel free to rename the models, but don't rename db_table values or field names.
from django.db import models


class Continent(models.Model):
    continentid = models.AutoField(primary_key=True)
    continentname = models.CharField(max_length=20, blank=True, null=True)

    class Meta:
        managed = False
        db_table = 'continents'


class Countrie(models.Model):
    countryid = models.AutoField(primary_key=True)
    countryname = models.CharField(max_length=100)
    continentid = models.ForeignKey('self', models.DO_NOTHING, db_column='continentid', blank=True, null=True)

    class Meta:
        managed = False
        db_table = 'countries'


class Podcast(models.Model):
    podcastid = models.AutoField(primary_key=True)
    podcastname = models.CharField(max_length=100)
    duration = models.CharField(max_length=20)
    podcastdescription = models.TextField()
    audioname = models.CharField(max_length=100)
    photoroute = models.CharField(blank=True, null=True)
    audioroute = models.CharField(blank=True, null=True)
    countryid = models.ForeignKey(Countrie, models.DO_NOTHING, db_column='countryid')

    class Meta:
        managed = False
        db_table = 'podcasts'
