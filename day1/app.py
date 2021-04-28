class cars:
    def __init__(self,Car_name,Car_size,Car_price):#constructor
        self.Car_name = Car_name
        self.Car_size = Car_size
        self.Car_price = Car_price
    def add_car(self):
        return "succesfully added car "+ self.Car_name
    def delete_car(self):
        return "succesfully deleted car "+ self.Car_name
    def reprice(self,new_price):
        old_price = self.Car_price
        self.Car_price = new_price
        return "repriced car "+self.Car_name + " from "+str(old_price)+" to "+str(self.Car_price)


new_car = cars("Tesla","55 kgs",5555)
print(new_car.reprice(22))
print(new_car.Car_price)
