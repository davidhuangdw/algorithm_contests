package google_code_jam.`2018`.`1b`

fun main() {
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull { if (it != "") it.toInt() else null }
    fun input_longs() = split_input().mapNotNull { if (it != "") it.toLong() else null }

    val (T) = input_ints()
    (1..T).forEach {
        val (M) = input_ints()
        val R = (1..M).map { input_ints().map { it - 1 } }
        val G = input_longs()
        val initSum = G.sum()
        fun set2(): Long {
            // todo: fix it - TLE for set2
            // trick: detect cycle without check call-stack: check the mono-increase value -- whether exceed the limit
            // mono-value: supply never change, 'need' recipe will increase every transform!
            var sum = 0L

            val rem = G.toMutableList()
            var need = hashMapOf(0 to 1L)
            while (need.values.sum() <= initSum) {
                var min = initSum
                for ((i, k) in need.toList()) {
                    min = minOf(min, rem[i] / k)  // bug: rem[i] not G[i]
                }

                sum += min
                // at least one will transform: so need.sum() will increase
                for ((i, k) in need.toList()) {
                    rem[i] -= min * k
                    if (rem[i] < k) {
                        need.remove(i)
                        for (x in R[i]) {
                            rem[x] += rem[i]
                            need[x] = need.getOrDefault(x, 0) + k
                        }
                        rem[i] = 0
                    }
                }
            }
            return sum
        }

        fun set3(): Long {
            var (lo, hi) = G[0] to initSum
            fun ok(t: Long): Boolean {
                val r = G.toMutableList()
                r[0] -= t
                for (_i in 0 until M) {
                    // trick: detect cycle: if cycle, will loop to self after M -> after M still have debts
                    var debt = false
                    for ((i, v) in r.withIndex()) {
                        if (v < 0) {
                            if (v < -initSum) return false
                            debt = true
                            r[i] = 0
                            for (x in R[i]) {
                                r[x] += v
                            }
                        }
                    }
                    if (!debt)
                        break
                }
                return r.all { it >= 0 }
            }
            while (lo <= hi) {
                val md = (lo + hi) / 2
                if (ok(md))
                    lo = md + 1
                else
                    hi = md - 1
            }
            return hi
        }

        val res = set3()
        println("Case #${it}: $res")
    }
}
