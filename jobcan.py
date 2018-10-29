#!/usr/bin/python3

import os
import datetime

from time import sleep
from selenium import webdriver
from selenium.webdriver.chrome.options import Options
from selenium.webdriver.common.keys import Keys
from datetime import datetime

# headless chrome
CHROME_DRIVER = os.path.expanduser('/usr/bin/chromedriver')
USER_NAME = 'xxxx@example.com'
PASSWORD = 'yyyyyy'

options = Options()
options.add_argument('--headless')
options.add_argument('--window-size=1280,3000')

driver = webdriver.Chrome(chrome_options=options)

driver.get("https://ssl.jobcan.jp/login/pc-employee/?client_id=your_company")
sleep(1)
driver.find_element_by_xpath("//*[@id='email']").send_keys(USER_NAME)
driver.find_element_by_xpath("//*[@id='password']").send_keys(PASSWORD + Keys.RETURN)
working_status_before_exec = driver.find_element_by_id("working_status").text
# print(working_status_before_exec)

## ここで乱数スリープするとボットにしてもバレにくくなるが、そんなことしてはダメだぞ
sleep(10)

## 打刻！
driver.find_element_by_id('adit-button-push').click()
sleep(2)

## result
working_status_after_exec = driver.find_element_by_id("working_status").text
### encode('utf-8')いるかも
#print(working_status_after_exec)
#print(working_status_after_exec.encode('utf-8'))

# snapshot
time_now = datetime.now().strftime("%Y%m%d%H%M%S")
file_name = time_now + "results_snapshot.png"
driver.save_screenshot(file_name)
