
from pynput.keyboard import Key,Listener 
from pygame import mixer
def on_release(key):
    if key==Key.esc:
    #Check to see if the key is the Escape key
        return False

def on_press(key):
    print("{0} pressed".format(key)) # Prints of key
    mixer.init() #initialize mixer
    mixer.music.load('audio.mp3') #load audio file
    mixer.music.play() #play the file

with Listener(on_press=on_press, on_release=on_release) as listener:
    listener.join()