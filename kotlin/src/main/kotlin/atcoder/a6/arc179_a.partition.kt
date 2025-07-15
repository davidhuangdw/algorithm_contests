package atcoder.a6

fun main(){
    fun inputInts() = readLine()!!.split(' ').map(String::toInt)

    val (N, K) = inputInts()
    val A = inputInts()
    if(0 < K){
        println("Yes")
        println(A.sorted().joinToString(" "))
        return
    }
    if(A.fold(0L,Long::plus) >= K){
        println("Yes")
        println(A.sortedDescending().joinToString(" "))
        return
    }
    println("No")
}
