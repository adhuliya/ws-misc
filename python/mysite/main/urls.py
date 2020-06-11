from django.urls import path

from . import views

#Ref: https://consideratecode.com/2018/05/02/django-2-0-url-to-path-cheatsheet/

app_name = "main"
urlpatterns = [
  path('', views.index, name='index'),
]


