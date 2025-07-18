package atcoder.a8

fun main() {
    fun inputInts() = readLine()!!.split(' ').map(String::toLong)
    val (N, A, B) = inputInts()
    println(if(N<A) 0 else if(A<=B) N+1-A else{
        val k = N/A
        k*B - maxOf(0, A*k+B-1 - N)
    })
}
