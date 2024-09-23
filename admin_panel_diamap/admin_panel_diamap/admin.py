from django.contrib import admin
from .models import Podcast,Continent,Countrie
class PodcastAdmin(admin.ModelAdmin):
    list_display = ('podcastid', 'podcastname', 'duration', 'audioname','photoroute', 'audioroute', 'countryid')
    fields = ('podcastname', 'duration', 'podcastdescription', 'audioname','photoroute', 'audioroute', 'countryid')
admin.site.register(Podcast,PodcastAdmin)
admin.site.register(Countrie)
admin.site.register(Continent)
