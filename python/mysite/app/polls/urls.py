from django.urls import path

from . import views

#Ref: https://consideratecode.com/2018/05/02/django-2-0-url-to-path-cheatsheet/

app_name = "polls"
urlpatterns = [
    # ex: /polls/
    path('', views.index, name='index'),
    # ex: /polls/5/
    path('<int:question_id>/', views.detail, name='detail'),
    # ex: /polls/5/results/
    path('<int:question_id>/results/', views.results, name='results'),
    # ex: /polls/5/vote/
    path('<int:question_id>/vote/', views.vote, name='vote'),
]


