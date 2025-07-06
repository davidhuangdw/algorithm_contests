package atcoder.a2

fun main() {
    readLine()!!
    val S = readLine()!!
    var sum = 0L
    var acc = 0L
    for(ch in S){
        if(ch == '>'){
            acc += 1
            sum += acc
        }else{
            acc = 0
        }
    }
    println(sum)
}