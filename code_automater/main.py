import json

class calculate:
    def __init__(self,filename):
        self.filename = filename
        
    def create_variable(self,variable,value):
        f = open(self.filename+".py","a")
        f.write(variable+"="+value)
        f.write("\n")
        return True

    def delete_variable(self,variable):
        #Text file removal
        with open(self.filename+".py","r") as file:
            lines = file.readlines()
        with open(self.filename+".py","w") as file:
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

    def show_var(self,variable):
        if calculate.check_var(self.filename,variable):
            string ="print("+variable+")"
            f = open(self.filename+".py","a")
            f.write("\n")
            f.write(string)
            f.write("\n")
            f.close()
        else:
            return "Error,var does not exist"

    def add_vars(self,new_var_name,*args):
        string = new_var_name+" = "
        added_vars = ""
        for i in args:
            if calculate.check_var(self.filename,i):
                added_vars+=i+"+"
            else:
                return "Error you variable is not found"
        added_vars = added_vars[0:-1]
        string += added_vars
        f = open(self.filename+".py","a")
        f.write("\n")
        f.write(string)
        f.write("\n")
        f.close()
    def add_import(self,module):
        string = "import "+module+"\n"
        with open(self.filename+".py") as f: 
            lines = f.readlines() #read 
            lines.insert(0,string)
            f.close()
        with open(self.filename+".py","w") as f:
            f.writelines(lines)
            f.close()
        return True

filename = input("Please enter name of your file: ")
while True:
    object = calculate(filename)
    print("------------------")
    print("Press 1 to create a variable")
    print("Press 2 to delete a variable")
    print("Press 3 to show a variable")
    print("Press 4 to add a library")
    print("Press 5 to add variables together ")
    print("Press anything else to quit")

    userinput = input()
    if userinput=="1":
        variable = input("What is the name of the variable: ")
        value = input("What is the value of the variable: ")
        if(object.create_variable(variable,value)):
            print("Succesfully done!")
        else:
            print("ERROR!!!!")
        
    break


#initialize = calculate("test")
#initialize.create_variable("Name","'Yassa taiseer'")
#initialize.create_variable("Name1","'Bruce wayne'")
#initialize.delete_variable("Name1")
#initialize.show_var("Name")
#initialize.add_import("pygame")




