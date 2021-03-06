from logging import logMultiprocessing
from manim import *
from numpy import square


class OpeningScene(Scene):
    def construct(self):
        image = VGroup()
        image +=  Text('Introduction To Linked Lists', color=BLUE,font_size=60, line_spacing=4)

        image += Text('By Yassa', font_size=35, line_spacing=4)
        
        self.add(image.arrange(DOWN).scale(.75))
class IndexScene(Scene):
    def construct(self):

        Line1 = Tex("Timeline ","Of ","This ", "Video ",color=BLUE,font_size=40).scale(3)
        Line1.move_to(UP*3)
        self.add(Line1)
        
        blist = BulletedList("What Are Linked Lists?", "Implementation", "Reversing A Linked List","Searching A Linked List",color=ORANGE, height=6, width=6)
        #blist.move_to(DOWN*2)
        
        self.add(blist)

class ExplainScene(Scene):
    def construct(self):

        Line1 = Tex("What ","is ","A ", "Linked ","List? ",color=BLUE,font_size=40).scale(3)

        Line1.move_to(UP*3)
        self.play(FadeIn(Line1, shift=UP, scale=0.66))
        blist = BulletedList("An Asbtract Data Strucuture", "Built on top of basic data structures(e.g int)", "Used to store data",color=ORANGE, height=12, width=12)
        blist.move_to(LEFT*1)
        self.play(FadeIn(blist, shift=DOWN, scale=0.66))
        self.play(Unwrite(Line1))
        self.play(Unwrite(blist))
class InfoScene(Scene):
    def construct(self):
        Node = Tex("Node:",font_size=15).scale(3)
        Node.move_to(UP*2+LEFT*5)

        Address = Tex("0x000",font_size=15).scale(3)
        Address.move_to(DOWN*2+LEFT*5)
        square1 = Square(side_length=2)
        square1.move_to(LEFT*5)
        Number1 = Tex("1",font_size=20).scale(3)
        Number1.move_to(LEFT*5)
        # some animations display mobjects, ...
        self.add(Number1)
        self.add(square1)
        self.add(Node)
        self.add(Address)
        self.wait(1)
        Pointer = Tex("---->",font_size=35).scale(3)
        Pointer.move_to(LEFT*2.5)
        self.play(FadeIn(Pointer, shift=UP, scale=0.66))
        square2 = Square(side_length=2)
        Number2 = Tex("2",font_size=20).scale(3)
        self.play(Write(square2))
        self.play(Write(Number2))
        self.wait(1)
        Pointer2 = Tex("---->",font_size=35).scale(3)
        Pointer2.move_to(RIGHT*3)
        self.play(FadeIn(Pointer2, shift=UP, scale=0.66))

        Number3 = Tex("NULL",font_size=20).scale(3)
        Number3.move_to(RIGHT*5)
        self.play(Write(Number3))
        self.wait(1)
        Head = Tex("head",font_size=14).scale(3)
        Head.move_to(DOWN*3+LEFT*5)
        self.play(Write(Head))
        Tail = Tex("tail",font_size=14).scale(3)
        Tail.move_to(DOWN*2)  
        self.play(Write(Tail))      
    # some animations remove mobjects from the screen
class ImplementScene(Scene):
    def construct(self):        
        Title = Tex("Implementation: ",font_size=18).scale(3)
        Title.move_to(LEFT*4+UP*3)
        self.add(Title)
        line1 = "class node("
        Structure = Tex(line1,color=BLUE,font_size=18).scale(3)
        Structure.move_to(LEFT*4+UP*1)
        self.play(Write(Structure))
        line2 = "int x"
        int1 = Tex(line2,color=BLUE,font_size=18).scale(3)
        int1.move_to(LEFT*4)
        self.play(Write(int1))

        line3 = "node* Next )"
        int2 = Tex(line3,color=BLUE,font_size=18).scale(3)
        int2.move_to(LEFT*3+DOWN*1)
        self.play(Write(int2))
        Pointer2 = Tex("---->",font_size=35).scale(3)
        Pointer2.move_to(DOWN*1)
        self.play(Write(Pointer2, shift=UP, scale=0.66))
        square1 = Square(side_length=2)
        square1.move_to(RIGHT*3+DOWN*1)
        Number1 = Tex("1",font_size=20).scale(3)
        Number1.move_to(RIGHT*3+DOWN*1)
        self.play(Write(square1))
        self.play(Write(Number1))
class linkedListShow(Scene):
    def construct(self):
        Title = Tex("Making A Linked List: ",font_size=14).scale(3)
        Title.move_to(LEFT*4+UP*3)
        self.play(Write(Title))
        self.play(Unwrite(Title))
        Node1 = Tex("Node FirstNode = new node()",color=BLUE,font_size=16).scale(3)
        NodeLine2 = Tex("FirstNode->value = 15",color=BLUE,font_size=16).scale(3)
        Node1.move_to(LEFT*3+UP*3)
        NodeLine2.move_to(LEFT*4+UP*2)
        self.play(Write(Node1))
        self.play(Write(NodeLine2))
        square1 = Square(side_length=2,color=BLUE)
        square1.move_to(RIGHT*1+DOWN*1)
        Number1 = Tex("15",font_size=20,color=BLUE).scale(3)
        Number1.move_to(RIGHT*1+DOWN*1)
        self.play(Write(square1)) 
        self.play(Write(Number1))
        Node2 = Tex("Node ScndNode = new node()",color=RED,font_size=16).scale(3)
        NodeLine3 = Tex("ScndNode->value = 16",color=RED,font_size=16).scale(3)
        Node2.move_to(LEFT*3+UP*1)
        NodeLine3.move_to(LEFT*4)
        self.play(Write(Node2))
        self.play(Write(NodeLine3))
        square2 = Square(side_length=2,color=RED)
        square2.move_to(RIGHT*5+DOWN*1)
        Number2 = Tex("16",font_size=20,color=RED).scale(3)
        Number2.move_to(RIGHT*5+DOWN*1)
        self.play(Write(square2))
        self.play(Write(Number2))
        NodeLine4 = Tex("FirstNode->next = ScndNode",color=ORANGE,font_size=16).scale(3)
        NodeLine4.move_to(LEFT*3.5+DOWN*1)
        self.play(Write(NodeLine4))
        Pointer = Tex("---->",color=ORANGE,font_size=14).scale(3)
        Pointer.move_to(RIGHT*3+DOWN*1)
        self.play(Write(Pointer))
class ReverseList(Scene):
    def construct(self):
        Title = Tex("Reversing A Linked List?",font_size=14).scale(3)
        Title.move_to(LEFT*4+UP*3)
        self.play(Write(Title))
        self.wait(1)
        self.play(Unwrite(Title))    
        square1 = Square(side_length=2,color=BLUE)
        square1.move_to(LEFT*6)
        self.play(Write(square1))
        Number1 = Tex("1",color=BLUE,font_size=14).scale(3)
        Number1.move_to(LEFT*6)
        self.play(Write(Number1))
        Pointer = Tex("---->",color=BLUE,font_size=24).scale(3)
        Pointer.move_to(LEFT*4)
        self.play(Write(Pointer))
        self.wait(1)

        square2 = Square(side_length=2,color=BLUE)
        square2.move_to(LEFT*2)
        self.play(Write(square2))
        Number2 = Tex("2",color=BLUE,font_size=14).scale(3)
        Number2.move_to(LEFT*2)
        Pointer2 = Tex("---->",color=BLUE,font_size=24).scale(3)
        self.play(Write(Pointer2))
        self.play(Write(Number2))
        self.wait(1)

        square3 = Square(side_length=2,color=BLUE)
        square3.move_to(RIGHT*2)
        self.play(Write(square3))
        Number3 = Tex("3",color=BLUE,font_size=14).scale(3)
        Number3.move_to(RIGHT*2)
        Pointer3 = Tex("---->",color=BLUE,font_size=24).scale(3)
        Pointer3.move_to(RIGHT*4)
        self.play(Write(Pointer3))
        self.play(Write(Number3))
        self.wait(1)

        Number4 = Tex("NULL",color=BLUE,font_size=14).scale(3)
        Number4.move_to(RIGHT*6)
        self.play(Write(Number4))
        self.wait(1)

        CurrentNode = Arrow(start=LEFT*6+DOWN*3,color=BLUE, end=LEFT*6)
        self.play(Write(CurrentNode))
        
        PrevNode = Arrow(start=LEFT*7+DOWN*3,color=RED, end=LEFT*7)
        self.play(Write(PrevNode))
        
        NextNode = Arrow(start=LEFT*5+DOWN*3,color=GREEN, end=LEFT*5)
        self.play(Write(NextNode))  

        square4 = Square(side_length=0.5,color=BLUE)
        square4.move_to(LEFT*6+UP*3)  
        square4.set_fill(BLUE,opacity=1.0)    
        Number5 = Tex("Current Node",color=BLUE,font_size=14).scale(3)
        Number5.move_to(LEFT*4+UP*3)   
        self.play(Write(square4))
        self.play(Write(Number5))
        square5 = Square(side_length=0.5,color=RED)
        square5.move_to(LEFT*2+UP*3)  
        square5.set_fill(RED,opacity=1.0)    
        Number6 = Tex("Prev. Node",color=RED,font_size=14).scale(3)
        Number6.move_to(UP*3) 
        self.play(Write(square5))
        self.play(Write(Number6))

        square6 = Square(side_length=0.5,color=GREEN)
        square6.move_to(RIGHT*2+UP*3)  
        square6.set_fill(GREEN,opacity=1.0)    
        Number7 = Tex("Next Node",color=GREEN,font_size=14).scale(3)
        Number7.move_to(RIGHT*4+UP*3) 
        self.play(Write(square6))
        self.play(Write(Number7))
        self.wait(1)
        CurrentNode.move_to(LEFT*2+DOWN*1.5)
        self.play(Write(CurrentNode))
        NextNode.move_to(DOWN*1.5)
        PrevNode.move_to(LEFT*4+DOWN*1.5)        
        
        self.wait(1)
        self.play(Unwrite(NextNode))
        self.play(Unwrite(PrevNode))
        PrevNode = Arrow(start=LEFT*2+DOWN*5 ,color=RED, end=LEFT*2+DOWN*1.5)
        NextNode = Arrow(start=LEFT*4+DOWN*5,color=GREEN, end=LEFT*4+DOWN*1.5)
        self.wait(1)    
        self.play(Write(NextNode))
        self.play(Write(PrevNode))
        self.wait(1)
        self.play(Unwrite(Pointer))
        Pointer = Tex("<----",color=YELLOW,font_size=24).scale(3)
        Pointer.move_to(LEFT*4)
        self.play(Write(Pointer))

class CodeReverse(Scene):
    def construct(self):
        Title = Tex("Code Demonstration: ",font_size=14).scale(3)
        Title.move_to(LEFT*4+UP*3)
        self.play(Write(Title))
        self.play(Unwrite(Title))
        self.wait(1)
        CurrentNode = Tex("  Node CurrentNode = headNode",color=BLUE,font_size=16).scale(3)
        CurrentNode.move_to(LEFT*3.5+UP*3)
        PrevNode = Tex("Node PrevNode = NULL",color=BLUE,font_size=16).scale(3)
        PrevNode.move_to(LEFT*4+UP*2)
        NextNode = Tex("Node NextNode = NULL",color=BLUE,font_size=16).scale(3)
        NextNode.move_to(LEFT*4+UP*1)
        self.play(Write(CurrentNode))
        self.play(Write(PrevNode))
        self.play(Write(NextNode))
        self.wait(1)
        Loop = Tex("while(CurrentNode!=NULL):",color=BLUE,font_size=16).scale(3)
        Loop.move_to(LEFT*3.5)
        self.play(Write(Loop))
        FinalCode1 = Tex("NextNode = CurrentNode->next;",font_size=11).scale(3)
        FinalCode1.move_to(LEFT*3+DOWN*1)
        FinalCode2 = Tex("CurrentNode->next = PrevNode;",font_size=11).scale(3)
        FinalCode2.move_to(LEFT*3+DOWN*1.5)
        FinalCode3 = Tex("PrevNode = CurrentNode;",font_size=11).scale(3)
        FinalCode3.move_to(LEFT*3+DOWN*2)
        FinalCode4 = Tex("CurrentNode = NextNode;",font_size=11).scale(3)
        FinalCode4.move_to(LEFT*3+DOWN*2.5)
        self.play(Write(FinalCode1))
        self.play(Write(FinalCode2))
        self.play(Write(FinalCode3))
        self.play(Write(FinalCode4))

class SearchList(Scene):
    def construct(self):
        Title = Tex("Searching A Linked List?",font_size=14).scale(3)
        Title.move_to(LEFT*4+UP*3)
        self.play(Write(Title))
        self.wait(1)
        self.play(Unwrite(Title))    
        square1 = Square(side_length=2,color=BLUE)
        square1.move_to(LEFT*6)
        self.play(Write(square1))
        Number1 = Tex("1",color=BLUE,font_size=14).scale(3)
        Number1.move_to(LEFT*6)
        self.play(Write(Number1))
        Pointer = Tex("---->",color=BLUE,font_size=24).scale(3)
        Pointer.move_to(LEFT*4)
        self.play(Write(Pointer))
        self.wait(1)

        square2 = Square(side_length=2,color=BLUE)
        square2.move_to(LEFT*2)
        self.play(Write(square2))
        Number2 = Tex("2",color=BLUE,font_size=14).scale(3)
        Number2.move_to(LEFT*2)
        Pointer2 = Tex("---->",color=BLUE,font_size=24).scale(3)
        self.play(Write(Pointer2))
        self.play(Write(Number2))
        self.wait(1)

        square3 = Square(side_length=2,color=BLUE)
        square3.move_to(RIGHT*2)
        self.play(Write(square3))
        Number3 = Tex("3",color=BLUE,font_size=14).scale(3)
        Number3.move_to(RIGHT*2)
        Pointer3 = Tex("---->",color=BLUE,font_size=24).scale(3)
        Pointer3.move_to(RIGHT*4)
        self.play(Write(Pointer3))
        self.play(Write(Number3))
        self.wait(1)

        Number4 = Tex("NULL",color=BLUE,font_size=14).scale(3)
        Number4.move_to(RIGHT*6)
        self.play(Write(Number4))
        self.wait(1)

        SearchNumber = Tex("Find Value: '2' ",color=ORANGE,font_size=20).scale(3)
        SearchNumber.move_to(LEFT*4+UP*3)
        self.play(Write(SearchNumber))


        CurrentNode = Arrow(start=LEFT*6+DOWN*3,color=BLUE, end=LEFT*6)
        self.play(Write(CurrentNode))
        self.wait(1)
        self.play(Unwrite(CurrentNode))

        CurrentNode = Arrow(start=LEFT*2+DOWN*3,color=BLUE, end=LEFT*2)
        self.play(Write(CurrentNode))
        self.wait(1)

        Complexity = Tex("Complexity: O(n)",color=WHITE,font_size=20).scale(3)
        Complexity.move_to(LEFT*4+DOWN*3)
        self.play(Write(Complexity))
  

class CodeSearch(Scene):
    def construct(self):
        Title = Tex("Code Demonstration: ",font_size=14).scale(3)
        Title.move_to(LEFT*4+UP*3)
        self.play(Write(Title))
        self.play(Unwrite(Title))
        self.wait(1)
        CurrentNode = Tex("  Node CurrentNode = headNode",color=BLUE,font_size=16).scale(3)
        CurrentNode.move_to(LEFT*3.5+UP*3)
        self.play(Write(CurrentNode))
        self.wait(1)

        Number = Tex("Number = 2",color=BLUE,font_size=16).scale(3)
        Number.move_to(LEFT*5.5+UP*2)
        self.play(Write(Number))
        self.wait(1)

        Loop = Tex("while(CurrentNode!=NULL):",color=BLUE,font_size=16).scale(3)
        Loop.move_to(LEFT*3.5)
        self.play(Write(Loop))
        FinalCode1 = Tex("if(CurrentNode->val==Number):",color=BLUE,font_size=16).scale(3)
        FinalCode1.move_to(LEFT*2+DOWN*1)
        FinalCode2 = Tex("return true",color=BLUE,font_size=16).scale(3)
        FinalCode2.move_to(LEFT*1+DOWN*1.5)
        FinalCode3 = Tex("CurrentNode=CurrentNode->next",color=BLUE,font_size=16).scale(3)
        FinalCode3.move_to(LEFT*2+DOWN*2)
        self.play(Write(FinalCode1))
        self.play(Write(FinalCode2))
        self.play(Write(FinalCode3))


class ThumbNail(Scene):
    def construct(self):
        Line1 = Tex("Linked List Crash Course ",color=BLUE,font_size=40).scale(3)
        Line1.move_to(UP*3)
        self.add(Line1)

        square1 = Square(side_length=2)
        square1.move_to(LEFT*5)
        Number1 = Tex("1",font_size=20).scale(3)
        Number1.move_to(LEFT*5)
        # some animations display mobjects, ...
        self.add(Number1)
        self.add(square1)
        Pointer = Tex("---->",font_size=35).scale(3)
        Pointer.move_to(LEFT*2.5)
        self.add(Pointer)
        square2 = Square(side_length=2)
        Number2 = Tex("2",font_size=20).scale(3)
        self.add(square2)
        self.add(Number2)
        Pointer2 = Tex("---->",font_size=35).scale(3)
        Pointer2.move_to(RIGHT*3)
        self.add(Pointer2)

        Number3 = Tex("NULL",font_size=20).scale(3)
        Number3.move_to(RIGHT*5)
        self.add(Number3)
        image = Text('By Yassa Taiseer', font_size=35, line_spacing=4)
        image.move_to(DOWN*2)
        self.add(image)