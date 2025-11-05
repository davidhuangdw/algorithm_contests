package atcoder.a10

fun main() {
    fun inputInts() = readLine()!!.split(' ').map(String::toInt)

    val (N) = inputInts()
    val sz = N*(N+1)/2
    val A = IntArray(sz){it+1}
    var i = N-2
    for(x in N downTo 2 step 2) {
        val y = x-1
        repeat(if(x == N) y else y-1){
            A[i++] = x
            A[i++] = y
        }
        A[i++] = x
    }
    println(A.joinToString(" "))
}
