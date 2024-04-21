from playwright.sync_api import sync_playwright
from selectolax.parser import HTMLParser


def get_html(page, asin):
    url = f"https://www.amazon.fr/dp/{asin}"
    page.goto(url)
    html = HTMLParser(page.content())
    return html


def parse_html(html, asin):
    print(html.css_first("title").text())
    print(asin)


def run():
    asin = "B08H93ZRK9"
    pw = sync_playwright().start()
    browser = pw.chromium.launch()
    page = browser.new_page()
    html = get_html(page, asin)
    parse_html(html, asin)


def main():
    run()


if __name__ == "__main__":
    main()
