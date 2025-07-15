package atcoder.a8

fun main(){
    fun inputInts() = readLine()!!.split(' ').map(String::toInt)

    readLine()
    val A = inputInts().sorted()
    val n = A.size
    var x = A[(n-1)/2]/2.0
    println("%.8f".format(A.map{maxOf(x, it-x)}.sum()/n))
}
