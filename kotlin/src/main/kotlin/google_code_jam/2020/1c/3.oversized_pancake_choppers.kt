package google_code_jam.`2020`.`1c`

fun main() {
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull { if (it != "") it.toInt() else null }
    fun input_longs() = split_input().mapNotNull { if (it != "") it.toLong() else null }

    val T = readLine()!!.toInt()
    (1..T).forEach {
        val (N, D) = input_ints()
        val A = input_longs().sortedDescending()

        fun set3(): Int { // todo: fix it
            // trick: condition_check outside -- precompute maxV
            var (l, r) = 1L to A.maxOrNull()!!
            while (l <= r) {
                val md = (l + r) / 2

                var c = 0L
                for (a in A) {
                    c += a / md
                }
                if (c >= D)
                    l = md + 1
                else
                    r = md - 1
            }
            val maxV = r

            val count = hashMapOf<Long, Int>()
            for (a in A)
                count[a] = count.getOrDefault(a, 0) + 1

            var maxSave = 1
            for (a in count.keys) {
                for (cut in 1 until D) {
                    if (!(a % cut == 0L && a / cut <= maxV))
                        continue

                    val v = a / cut
                    var r = D
                    var save = 0
                    for (k in 1..D) {
                        val c = minOf(r / k, count.getOrDefault(v * k, 0))
                        save += c
                        r -= c * k
                    }
                    maxSave = maxOf(maxSave, save)
                }
            }
            return D - maxSave
        }

        val res = set3()

        println("Case #${it}: $res")
    }
}
