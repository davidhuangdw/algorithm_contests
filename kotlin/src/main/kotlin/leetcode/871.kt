package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test
import java.util.*

class MinRefuelStops {
    fun minRefuelStops(target: Int, startFuel: Int, stations: Array<IntArray>): Int {
        val n = stations.size
        if (startFuel >= target) return 0

        val que = PriorityQueue<Int>(compareBy { -stations[it][1] })

        var (cur, cnt) = startFuel to 0
        var r = 0
        while (cur < target) {
            while (r < n && stations[r][0] <= cur) {
                que.add(r)
                r += 1
            }
            if (que.isEmpty()) break
            cnt += 1
            cur += stations[que.poll()][1]
        }

        return if (cur >= target) cnt else -1
    }

    @Test
    fun test1() {
        assertEquals(1, minRefuelStops(100, 50, parse2dIntArray("[[25,25],[50,50]]")))
        assertEquals(1, minRefuelStops(100, 50, parse2dIntArray("[[40,50]]")))
        assertEquals(2, minRefuelStops(100, 10, parse2dIntArray("[[10,60],[20,30],[30,30],[60,40]]")))
        assertEquals(0, minRefuelStops(1, 1, parse2dIntArray("[]")))
        assertEquals(-1, minRefuelStops(100, 1, parse2dIntArray("[[10,100]]")))
    }
}
