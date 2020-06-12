from django.urls import path

from . import views

#Ref: https://consideratecode.com/2018/05/02/django-2-0-url-to-path-cheatsheet/

app_name = "snippets"
urlpatterns = [
  path('', views.index, name='index'),
  path("visit_count/", views.visitCount, name='visitCount'),
  path('file/<path:filepath>', views.protectFile, name='protectFile'),
  path('jsonget/', views.jsonSend, name='jsonget'),
  path('jsonset/', views.jsonReceive, name='jsonset'),
  path('display_request/', views.displayRequest, name='displayRequest'),
]


