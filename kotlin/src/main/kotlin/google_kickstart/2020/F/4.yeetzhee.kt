package google_kickstart.`2020`.F

fun main() {
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().map { it.toInt() }
    fun input_longs() = split_input().map { it.toLong() }

    val T = readLine()!!.toInt()
    (1..T).forEach {
        val (n, m, k) = input_ints()
        val a = (1..k).map { input_ints()[0] }.sorted()

        val cache = hashMapOf<List<Int>, Double>()
        cache[a] = .0
        fun cal(state: List<Int>): Double {
            if (state !in cache) {
                val cur = state.toMutableList()
                var succ = 0
                var res = .0
                var i = k - 1
                while (i >= 0) {
                    if (state[i] == 0) break
                    cur[i] += 1
                    val st = cur.sorted()
                    if(st.zip(a).all{it.first <= it.second}){
                        succ += 1
                        res += cal(st)
                    }
                    cur[i] -= 1
                    i -= 1
                }
                if (i >= 0) {
                    val rem = m - (k - 1 - i)
                    cur[i] += 1
                    res += rem * cal(cur)
                    cur[i] -= 1
                    succ += rem
                }
                res = (res + m) / succ
                cache[state] = res
            }
            return cache[state]!!
        }

        val exp = cal((1..k).map { 0 })
        println("Case #${it}: $exp")
    }
}
