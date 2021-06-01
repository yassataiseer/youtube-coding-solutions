from pynput.keyboard import Key,Listener 

import time
from pygame import mixer
def on_release(key):
    if key==Key.esc:
        return False

def on_press(key):
    seconds = time.time()
    print("{0} pressed".format(key))
    mixer.init()
    mixer.music.load('audio.mp3')
    mixer.music.play()
    print(time.time()-seconds)
    


with Listener(on_press=on_press, on_release=on_release) as listener:
    listener.join()
