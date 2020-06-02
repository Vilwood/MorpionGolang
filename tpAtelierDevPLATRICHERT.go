package main
import "github.com/tadvi/winc"


    

func main() {
    mainWindow := winc.NewForm(nil)
    mainWindow.SetSize(400, 300)
    mainWindow.SetText("Jeu du morpion")


    btn0 := winc.NewPushButton(mainWindow)
    btn1 := winc.NewPushButton(mainWindow)
    btn2 := winc.NewPushButton(mainWindow)
    btn3 := winc.NewPushButton(mainWindow)
    btn4 := winc.NewPushButton(mainWindow)
    btn5 := winc.NewPushButton(mainWindow)
    btn6 := winc.NewPushButton(mainWindow)
    btn7 := winc.NewPushButton(mainWindow)
    btn8 := winc.NewPushButton(mainWindow)
    btnFermer := winc.NewPushButton(mainWindow)

    var Joueur string = "Joueur1" 

    

    btn0.SetText(".")
    btn0.SetPos(40, 50)
    btn0.SetSize(100, 40)
    btn0.OnClick().Bind(func(e *winc.Event){
        if (Joueur == "Joueur1"){
            btn0.SetText("X")
            Joueur = "Joueur2"
        }else if (Joueur == "Joueur2"){
            btn0.SetText("O")
            Joueur = "Joueur1"
        }
        //TestWin(btn0, btn1, btn2, btn3, btn4, btn5, btn6, btn7, btn8)
    })
    btn1.SetText(".")
    btn1.SetPos(40, 100)
    btn1.SetSize(100, 40)
    btn1.OnClick().Bind(func(e *winc.Event){
        if (Joueur == "Joueur1"){
            btn1.SetText("X")
            Joueur = "Joueur2"
        }else if (Joueur == "Joueur2"){
            btn1.SetText("O")
            Joueur = "Joueur1"
        }
        //TestWin(btn0, btn1, btn2, btn3, btn4, btn5, btn6, btn7, btn8)
    })
    btn2.SetText(".")
    btn2.SetPos(40, 150)
    btn2.SetSize(100, 40)
    btn2.OnClick().Bind(func(e *winc.Event){
        if (Joueur == "Joueur1"){
            btn2.SetText("X")
            Joueur = "Joueur2"
        }else if (Joueur == "Joueur2"){
            btn2.SetText("O")
            Joueur = "Joueur1"
        }
        //TestWin(btn0, btn1, btn2, btn3, btn4, btn5, btn6, btn7, btn8)
    })
    btn3.SetText(".")
    btn3.SetPos(140, 50)
    btn3.SetSize(100, 40)
    btn3.OnClick().Bind(func(e *winc.Event){
        if (Joueur == "Joueur1"){
            btn3.SetText("X")
            Joueur = "Joueur2"
        }else if (Joueur == "Joueur2"){
            btn3.SetText("O")
            Joueur = "Joueur1"
        }
        //TestWin(btn0, btn1, btn2, btn3, btn4, btn5, btn6, btn7, btn8)
    })
    btn4.SetText(".")
    btn4.SetPos(140, 100)
    btn4.SetSize(100, 40)
    btn4.OnClick().Bind(func(e *winc.Event){
        if (Joueur == "Joueur1"){
            btn4.SetText("X")
            Joueur = "Joueur2"
        }else if (Joueur == "Joueur2"){
            btn4.SetText("O")
            Joueur = "Joueur1"
        }
        //TestWin(btn0, btn1, btn2, btn3, btn4, btn5, btn6, btn7, btn8)
    })
    btn5.SetText(".")
    btn5.SetPos(140, 150)
    btn5.SetSize(100, 40)
    btn5.OnClick().Bind(func(e *winc.Event){
        if (Joueur == "Joueur1"){
            btn5.SetText("X")
            Joueur = "Joueur2"
        }else if (Joueur == "Joueur2"){
            btn5.SetText("O")
            Joueur = "Joueur1"
        }
        //TestWin(btn0, btn1, btn2, btn3, btn4, btn5, btn6, btn7, btn8)
    })
    btn6.SetText(".")
    btn6.SetPos(240, 50)
    btn6.SetSize(100, 40)
    btn6.OnClick().Bind(func(e *winc.Event){
        if (Joueur == "Joueur1"){
            btn6.SetText("X")
            Joueur = "Joueur2"
        }else if (Joueur == "Joueur2"){
            btn6.SetText("O")
            Joueur = "Joueur1"
        }
        //TestWin(btn0, btn1, btn2, btn3, btn4, btn5, btn6, btn7, btn8)
    })
    btn7.SetText(".")
    btn7.SetPos(240, 100)
    btn7.SetSize(100, 40)
    btn7.OnClick().Bind(func(e *winc.Event){
        if (Joueur == "Joueur1"){
            btn7.SetText("X")
            Joueur = "Joueur2"
        }else if (Joueur == "Joueur2"){
            btn7.SetText("O")
            Joueur = "Joueur1"
        }
        //TestWin(btn0, btn1, btn2, btn3, btn4, btn5, btn6, btn7, btn8)
    })
    btn8.SetText(".")
    btn8.SetPos(240, 150)
    btn8.SetSize(100, 40)
    btn8.OnClick().Bind(func(e *winc.Event){
        if (Joueur == "Joueur1"){
            btn8.SetText("X")
            Joueur = "Joueur2"
        }else if (Joueur == "Joueur2"){
            btn8.SetText("O")
            Joueur = "Joueur1"
        }
        //TestWin(btn0, btn1, btn2, btn3, btn4, btn5, btn6, btn7, btn8)
    })
    btnFermer.SetText("Fermer")
    btnFermer.SetPos(270, 200)
    btnFermer.SetSize(100, 40)
    btnFermer.OnClick().Bind(wndOnClose)
    mainWindow.Center()
    mainWindow.Show()
    mainWindow.OnClose().Bind(wndOnClose)

    winc.RunMainLoop()
}


func wndOnClose(arg *winc.Event) {
    winc.Exit()
}

/*func TestWin(btn0, btn1, btn2, btn3, btn4, btn5, btn6, btn7, btn8){
    var Win string = "."

    if (btn0.Text() !="." && btn0.Text()==btn3.Text()==btn6.Text()){
        Win=btn0.Text()
    }
    if (btn1.Text() !="." && btn1.Text()==btn4.Text()==btn7.Text()){
        Win=btn1.Text()
    }
    if (btn2.Text()!="." && btn2.Text()==btn5.Text()==btn8.Text()){
        Win=btn2.Text()
    }
    if (btn0.Text()!="." && btn0.Text()==btn1.Text()==btn2.Text()){
        Win=btn0.Text()
    }
    if (btn4.Text() !="." && btn4.Text()==btn5.Text()==btn3.Text()){
        Win=btn4.Text()
    }
    if (btn7.Text() !="." && btn7.Text()==btn8.Text()==btn6.Text()){
        Win=btn7.Text()
    }
    if (btn0.Text() !="." && btn0.Text()==btn4.Text()==btn8.Text()){
        Win=btn0.Text()
    }
    if (btn2.Text() !="." && btn2.Text()==btn4.Text()==btn6.Text()){
        Win=btn2.Text()
    }


    if (Win == "X"){
        Win = "Joueur 1 a gagné"
    }else if (Win == "O"){
        Win = "Joueur 2 a gagné"
    }else{
        Win = "Continuez de jouer !"
    }
    return (Win)

}*/
