package atcoder.a4

fun main() {
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull { if (it != "") it.toInt() else null }

    val (N) = input_ints()
    val A = readLine()!!.reversed().map { it.digitToInt() }.toMutableList()
    val B = readLine()!!.reversed().map { it.digitToInt() }.toMutableList()

    val MOD = 998244353L
    val pow = (1..N).scan(1L){ acc, _ -> (acc * 10) % MOD}

    for (i in 0 until N) {
        if (A[i] > B[i]) {
            A[i] = B[i].also { B[i] = A[i] }
        }
    }

    var ab = listOf(A, B).map{it.foldIndexed(0L){ i, acc, v -> (acc + v * pow[i]) % MOD }}.reduce{a, b -> (a*b) % MOD}
    println(ab)
}
