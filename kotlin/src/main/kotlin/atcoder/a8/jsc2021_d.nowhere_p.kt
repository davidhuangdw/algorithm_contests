package atcoder.a8

fun main() {
    fun inputInts() = readLine()!!.split(' ').map(String::toInt)

    val MOD = 1e9.toInt() + 7
    val (N, P) = inputInts()

    var cur = P - 1L
    var r = N-1
    var b = P-2L // bug: long

    while(r > 0){
        if(r and 1 > 0){
            cur = (cur * b) % MOD
        }
        r /= 2
        b = b * b % MOD
    }
    println(cur)
}
