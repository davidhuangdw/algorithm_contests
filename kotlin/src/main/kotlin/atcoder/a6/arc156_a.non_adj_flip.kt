package atcoder.a6

fun main() {
    fun inputInts() = readLine()!!.split(' ').map(String::toInt)

    fun sol(s: String): Int {
        val n = s.length
        val idx = (0 until n).filter {s[it] == '1'}
        val k = idx.size
        if(k and 1 == 1) return -1
        if(k == 0 || k >= 4 || idx[1] - idx[0] > 1) return k/2
        if(n <= 3) return -1
        if(idx[0] >= 2 || n-1-idx[1] >= 2) return 2
        return 3
    }

    val (T) = inputInts()
    repeat(T) {
        readLine()
        println(sol(readLine()!!))
    }
}
