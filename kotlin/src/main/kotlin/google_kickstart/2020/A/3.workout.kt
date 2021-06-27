package google_kickstart.`2020`.A

fun main() {
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull { if (it != "") it.toInt() else null }
    fun input_longs() = split_input().mapNotNull { if (it != "") it.toLong() else null }

    val T = readLine()!!.toInt()
    (1..T).forEach {
        val (N, K) = input_ints()
        val A = input_ints()
        val D = A.zipWithNext { a, b -> b - a }

        fun valid(d: Int): Boolean {
            var cnt = 0
            for (diff in D) {
                cnt += (diff - 1) / d       // ceil(diff/d) - 1
                if (cnt > K) return false
            }
            return true
        }

        var (l, r) = 1 to A.last()
        while (l <= r) {
            val md = (l + r) ushr 1
            if (valid(md)) r = md - 1
            else l = md + 1
        }
        println("Case #${it}: $l")
    }
}
