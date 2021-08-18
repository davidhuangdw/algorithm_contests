package google_kickstart.`2021`.acpc

fun main() {
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull { if (it != "") it.toInt() else null }
    fun input_longs() = split_input().mapNotNull { if (it != "") it.toLong() else null }

    val T = readLine()!!.toInt()
    (1..T).forEach {
        val (E, N) = input_ints()
        val A = input_longs().sorted()
        var r = N - 1
        var mx = 0
        var e = E.toLong()
        var h = 0

        for ((i, s) in A.withIndex()) {
            if (i > r) break
            if (e <= s && h > 0) {    // bug: shouldn't skip i after acquire the largest team
                h -= 1
                e += A[r]
                r -= 1
            }
            if (i > r) break
            if (e > s) {
                h += 1
                e -= s
                mx = maxOf(mx, h)
            } else break
        }
        println("Case #${it}: $mx")
    }
}
