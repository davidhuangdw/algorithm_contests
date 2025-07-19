package atcoder.a8

fun main() {
    fun inputInts() = readLine()!!.split(' ').map(String::toInt)
    val x = inputInts().toHashSet()

    val (N) = inputInts()
    val (A, B) = List(2){inputInts()}
    val cnt = B.groupBy { it }.mapValues { it.value.size }

    fun check(x: Int): Boolean {
        val c = hashMapOf<Int, Int>()
        return A.all{a ->
            val b = a xor x
            val k = (c[b] ?: 0) + 1
            c[b] = k
            k <= (cnt[b] ?:0)
        }
    }

    val res = B.map{it xor A[0]}.toHashSet().filter{check(it)}.sorted()
    println(res.size)
    res.forEach { println(it) }
}
