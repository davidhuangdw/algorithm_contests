package atcoder.a5

fun main(){
    fun input_ints() = readLine()!!.split(" ").map { it.toInt() }

    val (T) = input_ints()
    repeat(T) {
        val (N) = input_ints()
        val P = input_ints()
        var mn = N+1
        println(P.reversed().count { v ->
            if((mn < v).also{mn = minOf(v, mn)}) false else true
        })
    }
}
