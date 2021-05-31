import json

def create_variable(filename,variable,value):
    f = open(filename+".py","a")
    f.write(variable+"="+value)
    f.write("\n")
    return True

def delete_variable(filename,variable):
    #Text file removal
    with open(filename+".py","r") as file:
        lines = file.readlines()
    with open(filename+".py","w") as file:
        for line in lines:
            line = line.strip("\n")
            if variable not in line: 
                file.write(line)

def check_var(filename,variable):
    counter = 0
    with open(filename+".py","r") as file:
        lines = file.readlines()
    for line in lines:
        if counter >0:
            return True
        line = line.strip("\n")
        if variable in line:
            counter +=1
    if counter>0:
        return True
    else:
        return False

def show_var(filename,variable):
    if check_var(filename,variable):
        string ="print("+variable+")"
        f = open(filename+".py","a")
        f.write("\n")
        f.write(string)
        f.write("\n")
        f.close()
    else:
        return "Error,var does not exist"

def add_vars(filename,new_var_name,*args):
    string = new_var_name+" = "
    added_vars = ""
    for i in args:
        if check_var(filename,i):
            added_vars+=i+"+"
        else:
            return "Error you variable is not found"
    added_vars = added_vars[0:-1]
    string += added_vars
    f = open(filename+".py","a")
    f.write("\n")
    f.write(string)
    f.write("\n")
    f.close()
def add_import(filename,module):
    string = "import "+module+"\n"
    with open(filename+".py") as f: 
        lines = f.readlines() #read 
        lines.insert(0,string)
        f.close()
    with open(filename+".py","w") as f:
        f.writelines(lines)
        f.close()
    return True
#add_import("test","pygame")
#add_vars("test","new_var","Name","Other_var")
#show_var("test","new_var")
#print(check_var("test","Na4231me101"))
#delete_variable("test","Name")