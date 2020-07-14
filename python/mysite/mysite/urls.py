"""mysite URL Configuration

The `urlpatterns` list routes URLs to views. For more information please see:
    https://docs.djangoproject.com/en/3.0/topics/http/urls/
Examples:
Function views
    1. Add an import:  from my_app import views
    2. Add a URL to urlpatterns:  path('', views.home, name='home')
Class-based views
    1. Add an import:  from other_app.views import Home
    2. Add a URL to urlpatterns:  path('', Home.as_view(), name='home')
Including another URLconf
    1. Import the include() function: from django.urls import include, path
    2. Add a URL to urlpatterns:  path('blog/', include('blog.urls'))
"""
from django.contrib import admin
from django.urls import path, include

# The bad_request() view is overridden by handler400:
handler400 = 'main.views.handler000'

# The permission_denied() view is overridden by handler403:
handler403 = 'main.views.handler403'

# The page_not_found() view is overridden by handler404:
handler404 = 'main.views.handler404'

# The server_error() view is overridden by handler500:
handler500 = 'main.views.handler500'

urlpatterns = [
  path('', include('main.urls')),  # the 'main' app
  path('snippets/', include('snippets.urls')),
  path('polls/', include('polls.urls')),
  path('admin/', admin.site.urls),
]
