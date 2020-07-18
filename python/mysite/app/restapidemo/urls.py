from django.urls import path

from . import views

#Ref: https://consideratecode.com/2018/05/02/django-2-0-url-to-path-cheatsheet/
from rest_framework import routers
from . import views

router = routers.DefaultRouter()
router.register(r'users', views.UserViewSet)
router.register(r'groups', views.GroupViewSet)
# include router.urls in mysite.urls

app_name = "restapidemo"
# urlpatterns = [
#   path('users/', views.UserViewSet.as_view, name='users'),
#   path('groups/', views.GroupViewSet.as_view, name='groups'),
# ]


