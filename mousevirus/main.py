import pyautogui
import time

while True:
    pyautogui.moveTo(200,100,duration=0.5)
    time.sleep(5)
    pyautogui.press('a')
