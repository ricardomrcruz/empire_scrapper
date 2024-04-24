from django.contrib import admin
from django.urls import path, include

from playground.views import index


# hello
urlpatterns = [
    path("", index, name="index"),
    path("admin/", admin.site.urls),
]
