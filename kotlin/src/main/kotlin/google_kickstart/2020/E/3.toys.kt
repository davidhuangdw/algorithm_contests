package google_kickstart.`2020`.E

import java.util.*

fun main() {
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().map { it.toInt() }
    fun input_longs() = split_input().map { it.toLong() }

    val T = readLine()!!.toInt()
    (1..T).forEach {
        val (n) = input_ints()
        val E = LongArray(n)
        val R = LongArray(n)
        for (i in 0 until n) {
            val x = input_longs()
            E[i] = x[0]
            R[i] = x[1]
        }

        var sum = E.sum()
        var (cur, cnt) = sum to 0
        var ma = cur to cnt
        val pq = PriorityQueue<Pair<Long, Int>>(compareBy { -it.first })

        for (i in 0 until n) {
            // to pass i not cry:
            if (E[i] + R[i] <= sum) {
                cur += E[i]
                if (cur > ma.first)
                    ma = cur to cnt
                pq.add(E[i] + R[i] to i)
            } else {
                sum -= E[i]
                cur -= E[i]         // bug: forgot -E[i]
                cnt += 1
                while (pq.isNotEmpty() && pq.peek().first > sum) {
                    val j = pq.poll().second
                    sum -= E[j]
                    cur -= 2 * E[j]
                    cnt += 1
                }
            }
        }
        val res = if (pq.size > 0) "$cnt INDEFINITELY" else "${ma.second} ${ma.first}"
        println("Case #${it}: $res")
    }
}
