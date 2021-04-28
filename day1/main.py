class cars:
    def __init__(self,Car_Name,Car_size,Car_price):
        self.Car_Name = Car_Name
        self.Car_Size = Car_size
        self.Car_price = Car_price
    def add_db(self):
        return "Succesfully added to the database " + self.Car_Name +" at the price of "+ str(self.Car_price)
    def delete_car(self):
        return "Deleted from database!"
    def reprice(self,new_price):
        old_price = self.Car_price
        self.Car_price = new_price
        return "repriced the car from "+str(old_price)+" to "+str(self.Car_price) 

new_car = cars("Tesla","1000 kg",100000)
print(new_car.reprice(5000))
print(new_car.Car_price)
