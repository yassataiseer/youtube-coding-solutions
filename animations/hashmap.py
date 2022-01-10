from logging import logMultiprocessing
from manim import *
from numpy import square

## manim -pql scene.py {{Current Scene}}

class OpeningScene(Scene):
    def construct(self):
        image = VGroup()
        image +=  Text('Why Are HashMaps So Efficient', color=BLUE,font_size=60, line_spacing=4)

        image += Text('By Yassa', font_size=35, line_spacing=4)
        
        self.add(image.arrange(DOWN).scale(.75))

class Efficiency(Scene):
    def construct(self):
        t2 = Table(
            [["~O(1)"],["~O(1)"],["O(1)"]],
            row_labels=[Text("Search"), Text("Delete"),Text("Insert")],
            col_labels=[Text("TIME COMP.")],
            top_left_entry=Text("Actions."),
            include_outer_lines=True,
            arrange_in_grid_config={"cell_alignment": RIGHT})
        self.play(Write(t2))

class PreReqs(Scene):
    def construct(self):
        Line1 = Tex("Prerequisites",color=BLUE,font_size=40).scale(3)
        Line1.move_to(UP*3)
        self.add(Line1)
        
        blist = BulletedList("Memory management", "Linked List","High School Math(ish)","Basic Knowledge of Hashmaps",color=ORANGE, height=6, width=6)
        #blist.move_to(DOWN*2)
        
        self.add(blist)

class Array(Scene):
    def construct(self):
        m0 = IntegerMatrix(
            [[5, 2,1],],
            left_bracket="(",
            right_bracket=")")
        m0.move_to(RIGHT*2)
        self.play(Write(m0))
        self.wait(1)

        code = Tex("print(Array[1])")
        code.move_to(LEFT*4)
        self.play(Write(code))
        self.wait(1)

        Pointer = Arrow(start=RIGHT*2+DOWN*2,color=WHITE, end=RIGHT*2)
        self.play(Write(Pointer))
        Answer = Tex("Answer: 2")
        Answer.move_to(RIGHT*2+DOWN*3)
        self.play(Write(Answer))
        TimeC = Tex("Time Comp: O(1)")
        TimeC.move_to(RIGHT*2+DOWN*3.5)
        self.play(Write(TimeC))

class MapStuff1(Scene):
    def construct(self):
        Number2 = Tex("Key",color=BLUE,font_size=14).scale(3)
        Number2.move_to(LEFT*2)
        Pointer2 = Tex("---->",color=BLUE,font_size=24).scale(3)
        Value = Tex("Value",color=BLUE,font_size=28).scale(3)
        Value.move_to(RIGHT*2)
        self.play(Write(Pointer2))
        self.play(Write(Number2))
        self.play(Write(Value))
        self.wait(1)

        self.play(Unwrite(Number2))
        self.play(Unwrite(Value))
        Number2 = Tex("Yassa Taiseer",color=BLUE,font_size=14).scale(3)
        Number2.move_to(LEFT*2.5)
        Value = Tex("16",color=BLUE,font_size=28).scale(3)
        Value.move_to(RIGHT*2)
        self.play(Write(Number2))
        self.play(Write(Value))

class MapStuff2(Scene):
    def construct(self):
        image = VGroup()

        image +=  Square(side_length=2).move_to(LEFT*6)
        image +=  Square(side_length=2).move_to(LEFT*5)
        image +=  Square(side_length=2).move_to(LEFT*4)
        image +=  Square(side_length=2).move_to(LEFT*3)
        image +=  Square(side_length=2).move_to(LEFT*2)
        image +=  Square(side_length=2).move_to(LEFT*1)
        image +=  Square(side_length=2)
        image +=  Square(side_length=2).move_to(RIGHT*1)
        image +=  Square(side_length=2).move_to(RIGHT*2)
        self.play(Write(image))
        self.wait(1)

        Value = Tex(" 'Yassa'--> 200",color=BLUE,font_size=18).scale(3)
        Value.move_to(LEFT*5+UP*3)
        self.play(Write(Value))
        self.wait(1)


        Value1 = Tex(" H(200)--> M==42 ",color=BLUE,font_size=18).scale(3)
        Value1.move_to(LEFT*4+UP*2)
        self.play(Write(Value1))
        self.wait(1)


        Value2 = Tex(" 42 mod array-size = 2",color=BLUE,font_size=18).scale(3)
        Value2.move_to(RIGHT*2+UP*3)
        self.play(Write(Value2))

        self.wait(1)
        Pointer = Arrow(start=LEFT*4.5+DOWN*2,color=WHITE, end=LEFT*4.5)
        self.play(Write(Pointer))



class MapStuff3(Scene):
    def construct(self):
        Value2 = Tex(" 42 mod array-size = 2",color=BLUE,font_size=18).scale(3)
        Value2.move_to(LEFT*4+UP*3)
        self.play(Write(Value2))

        Value3 = Tex(" 52 mod array-size = 2",color=BLUE,font_size=18).scale(3)
        Value3.move_to(LEFT*4+UP*2)
        self.play(Write(Value3))
        self.wait(1)

        image = VGroup()

        image +=  Square(side_length=2).move_to(LEFT*6)
        image +=  Square(side_length=2).move_to(DOWN*1+LEFT*6)
        image +=  Square(side_length=2).move_to(DOWN*2+LEFT*6)
        image +=  Square(side_length=2).move_to(DOWN*3+LEFT*6)

        self.play(Write(image))
        self.wait(1)

        Pointer = Arrow(start=DOWN*1.5+LEFT*6,color=WHITE, end=DOWN*1.5+LEFT*3)
        self.play(Write(Pointer))

        image1 = VGroup()
        image1 +=  Square(side_length=2).move_to(DOWN*1.5+LEFT*3)
        image1 += Tex("42",color=BLUE,font_size=18).scale(3).move_to(DOWN*1.5+LEFT*3)
        self.play(Write(image1))


        Pointer1 = Arrow(start=DOWN*1.5+LEFT*3,color=WHITE, end=DOWN*1.5+RIGHT*1)
        self.play(Write(Pointer1))

        image2 = VGroup()
        image2 +=  Square(side_length=2).move_to(DOWN*1.5+RIGHT*1)
        image2 += Tex("52",color=BLUE,font_size=18).scale(3).move_to(DOWN*1.5+RIGHT*1)
        self.play(Write(image2))
