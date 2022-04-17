package google_code_jam.`2020`.`1b`

fun main() {
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull { if (it != "") it.toInt() else null }
    fun input_longs() = split_input().mapNotNull { if (it != "") it.toLong() else null }

    class Succ : RuntimeException("Succ")
    class Fail : RuntimeException("Fail")
    val (T, A, B) = input_ints()
    val H = 1000_000_000L
    val CENTER = "CENTER"
    val HIT = "HIT"
    val MISS = "MISS"
    val WRONG = "WRONG"
    fun match(p: LongArray, target: String): Boolean {
        println(p.joinToString(" "))
        val (resp) = split_input()
        if (resp == CENTER)
            throw Succ()
        if (resp == WRONG)
            throw Fail()
        return resp == target
    }

    fun hit(p: LongArray) = match(p, HIT)
    fun findHit(): LongArray {
        val d = H / 2
        var x = -H + d
        for (i in 0 until 3) {
            var y = -H + d
            for (j in 0 until 3) {
                val p = longArrayOf(x, y)
                if (hit(p))
                    return p
                y += d
            }
            x += d
        }
        throw RuntimeException("fail to hit")
    }
    for (_t in 1..T) {
        try {
            val hit = findHit()
            val poss = List(2) { mutableListOf<Long>() }
            for ((i, v) in hit.withIndex()) {
                var cur = hit.clone()
                // +a
                var (l, r) = 0L to H - v
                while (l <= r) {
                    val md = (l + r) / 2
                    cur[i] = v + md
                    if (hit(cur)) l = md + 1
                    else r = md - 1
                }
                val a = r

                // -b
                l = 0L
                r = v + H
                while (l <= r) {
                    val md = (l + r) / 2
                    cur[i] = v - md
                    if (hit(cur)) l = md + 1
                    else r = md - 1
                }
                val b = r

                poss[i].add(v + (a - b) / 2)
                if ((a - b) % 2 != 0L) {
                    poss[i].add(v + (a - b + 1) / 2)
                }
            }

            for (x in poss[0])
                for (y in poss[1]) {
                    hit(longArrayOf(x, y))
                }
        } catch (e: Succ) {
        }
//        println("Case #${it}: ")
    }
}
