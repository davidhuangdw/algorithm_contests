package google_code_jam.`2019`.qua

import java.math.BigInteger

fun main() {
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull { if (it != "") it.toInt() else null }
    fun input_longs() = split_input().mapNotNull { if (it != "") it.toLong() else null }

    fun gcd(_x: BigInteger, _y: BigInteger): BigInteger {
        var (x, y) = _x to _y
        while (x != BigInteger("0")) {
            y = x.also { x = y % x }
        }
        return y
    }

    val T = readLine()!!.toInt()
    (1..T).forEach {
        val L = split_input().last().toInt()  // bug: N is bigInt
        fun set1(): String {
            val A = split_input().map { BigInteger(it) }
            val p = MutableList(L + 1) { BigInteger("0") }
            val k = (0 until L).firstOrNull { A[it] != A[it + 1] }!!  // bug: the case ABABABC...
            p[k + 1] = gcd(A[k], A[k + 1])
            for (i in k downTo 0) {
                p[i] = A[i] / p[i + 1]
            }
            for (i in k + 2..L) {
                p[i] = A[i - 1] / p[i - 1]
            }
            val mp = hashMapOf<BigInteger, Char>()
            for ((i, v) in p.toHashSet().sorted().withIndex()) {
                mp[v] = 'A' + i
            }
            return p.map { mp[it] }.joinToString("")
        }

        val res = set1()
        println("Case #${it}: $res")
    }
}
