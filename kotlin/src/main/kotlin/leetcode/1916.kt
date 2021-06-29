package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test

class WaysToBuildRooms {
    fun waysToBuildRooms(prevRoom: IntArray): Int {
        val MOD = (1e9 + 7).toLong()
        val n = prevRoom.size
        val ch = List(n) { mutableListOf<Int>() }
        for (i in 1 until n) {
            ch[prevRoom[i]].add(i)
        }

        val fact = MutableList(n + 1) { 1L }
        val inv = MutableList(n + 1) { 1L }
        val inv_fact = MutableList(n + 1) { 1L }
        for (i in 2..n) {
            fact[i] = fact[i - 1] * i % MOD
            inv[i] = (MOD - MOD / i * inv[(MOD % i).toInt()] % MOD) % MOD
            inv_fact[i] = inv_fact[i - 1] * inv[i] % MOD
        }

        fun calc(u: Int): Pair<Int, Long> {
            var res = 1L
            var sz = 1

            for (v in ch[u]) {
                val (ch_sz, ch_res) = calc(v)
                res = res * ch_res % MOD
                res = res * inv_fact[ch_sz] % MOD
                sz += ch_sz
            }
            res = res * fact[sz - 1] % MOD
            return sz to res
        }

        return calc(0).second.toInt()
    }

    @Test
    fun testWaysToBuildRooms() {
        assertEquals(3, waysToBuildRooms(listOf(-1, 0, 1, 2, 1).toIntArray()))
    }

}