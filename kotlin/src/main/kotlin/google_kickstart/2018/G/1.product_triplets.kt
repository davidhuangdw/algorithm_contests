fun main() {
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull { if (it != "") it.toInt() else null }
    fun input_longs() = split_input().mapNotNull { if (it != "") it.toLong() else null }

    val T = readLine()!!.toInt()
    (1..T).forEach { // a * b == z
        val (N) = input_ints()
        val A = input_ints()
        val count = hashMapOf<Int, Int>()
        for (v in A)
            count[v] = count.getOrDefault(v, 0) + 1
        var cnt = 0L
        for (v in count.keys.sorted()) {
            val c = count[v]!!
            if (v == 0) {
                cnt += 1L * c * (c - 1) * (c - 2) / 6 + 1L * c * (c - 1) / 2 * (N - c)
            } else if (v == 1) {
                cnt += 1L * c * (c - 1) * (c - 2) / 6
                for ((x, cx) in count)
                    if (x > 1)
                        cnt += 1L * c * cx * (cx - 1) / 2
            } else {
                for (a in 2 until v) {
                    if (a * a > v) break
                    val b = v / a
                    if (a in count && v % a == 0 && b in count) {
                        val (ca, cb) = count[a]!! to count[b]!!
                        cnt += 1L * c * (if (a == b) ca * (ca - 1) / 2 else ca * cb)
                    }
                }
            }
        }
        println("Case #${it}: $cnt")
    }
}
