package main
import(
"fmt"
"os"
"os/exec")

func main(){
    var dicerolls [16]int16
    total := 0

    var a,b,c,d int16 
    for a,b,c,d = 1,1,1,1; a < 7 
    {
        total++
        dicerolls[rollCal(a,b,c,d) -3] = dicerolls[rollCal(a,b,c,d) -3] +1
        d++
        if d > 6 { d, c = 1, c +1}
        if c > 6 { c, b = 1, b +1}
        if b > 6 { b, a = 1, a +1}}

    graphArray(dicerolls[:], 30, 6, 6)
    printrolls(dicerolls[:], total)
}
func sort(arr []int16) {
    var i int = 1
    for i < len(arr){
        var j int = i
        for j > 0 && arr[j-1] > arr[j]{
            arr[j], arr[j-1] = arr[j-1], arr[j]
            j = j-1}
        i = i+1}
}
func rollCal( a int16, b int16, c int16, d int16,) int16 {
    arr := [4]int16{a,b,c,d}
    sort(arr[:])
    sum := arr[1] + arr[2] + arr[3]
    return sum
}
func printrolls( arr []int16, total int){
    for i := 0; i < len(arr); i++ {
	percent := (100*(float32(arr[i])/float32(total)))
        fmt.Printf("there is a %v in %v chance to roll a %v = %v%%. \n", arr[i], total, i +3, percent)
}

}
func graphArray(arr []int16, row int16, Xscrunch int16, Ystrech int){
    c := exec.Command("clear")
    c.Stdout = os.Stdout
    c.Run()
    var i int16
    for i = row ; i > 0
    {
        fmt.Printf("|")
        for j := 0; j < len(arr); j++{
            var chr string 
            if arr[j] > i*Xscrunch{
                chr = "X"
            }else{ chr = " "}
            for x := 1; x <= Ystrech; x++{
                fmt.Printf("%v", chr)}
        }
        fmt.Println("|")
        i = i -1
    }
}
