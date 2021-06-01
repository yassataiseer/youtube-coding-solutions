from pynput.keyboard import Key,Listener 

from playsound import playsound


def on_release(key):
    if key==Key.esc:
        return False


def on_press(key):
    print("{0} pressed".format(key))
    playsound('audio.mp3')
    


with Listener(on_press=on_press, on_release=on_release) as listener:
    listener.join()
