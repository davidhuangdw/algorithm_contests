package google_kickstart.`2021`.acpc

fun main() {
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull { if (it != "") it.toInt() else null }
    fun input_longs() = split_input().mapNotNull { if (it != "") it.toLong() else null }

    val T = readLine()!!.toInt()
    (1..T).forEach {
        val (N, K) = input_ints()
        val A = input_ints()

        var cnt = 0
        var i = 0
        while (i < N) {
            if (A[i] != K)
                i += 1
            else {
                var j = i
                while (j < N && A[j] == K - (j - i)) {
                    j += 1
                }
                if (j - i >= K)
                    cnt += 1
                i = j
            }
        }
        println("Case #${it}: $cnt")
    }
}
