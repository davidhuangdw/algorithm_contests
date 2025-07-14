package atcoder.a7

fun main(){
    fun inputInts() = readLine()!!.split(" ").map { it.toInt() }

    fun succ(n: Int, a: Int, b: Int): Boolean {
        if(a > n) return false
        val r = n - a
        return b <= r * minOf((n+1)/2, r)
    }

    val (T) = inputInts()
    repeat(T) {
        val (N, A, B) = inputInts()
        println(if(succ(N, A, B)) "Yes" else "No")

    }
}
