package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test
import java.util.*

class AssignTasks1882 {
    fun assignTasks(servers: IntArray, tasks: IntArray): IntArray {
        data class Point(val t: Int, val start: Boolean, val i: Int) : Comparable<Point> {
            override fun compareTo(other: Point) = t - other.t
        }

        val assign = IntArray(tasks.size) { -1 }

        val points = PriorityQueue<Point>()
        val S = PriorityQueue<Int>(compareBy({ servers[it] }, { it }))
        val wait_tasks = PriorityQueue<Int>(compareBy { it })

        for (i in tasks.indices) points.add(Point(i, true, i))
        for (i in servers.indices) S.add(i)
        while (points.isNotEmpty()) {
            val t = points.peek().t
            while (points.isNotEmpty() && points.peek().t == t) {
                val p = points.poll()
                if (p.start) {
                    wait_tasks.add(p.i)
                } else {
                    S.add(p.i)
                }
            }

//            println("\n-------time: $t: wait: ${wait_tasks.toList()} servers: ${S.toList().sortedWith(compareBy({ servers[it] }, { it }))}")
            while (S.isNotEmpty() && wait_tasks.isNotEmpty()) {
                val task = wait_tasks.poll()
                val s = S.poll()
                assign[task] = s
//                println("assign: $t: task $task to server $s release at ${t+tasks[task]}")
                points.add(Point(t + tasks[task], false, s))
            }
//            println("-------after: $t: wait: ${wait_tasks.toList()} servers: ${S.toList().sortedWith(compareBy({ servers[it] }, { it }))}")
        }

        return assign
    }

    @Test
    fun test1() {
        assertEquals(
            parseIntArray("[8,0,3,9,5,1,10,6,4,2,7,9,0]").toList(),
            assignTasks(
                parseIntArray("[10,63,95,16,85,57,83,95,6,29,71]"),
                parseIntArray("[70,31,83,15,32,67,98,65,56,48,38,90,5]")
            ).toList()
        )
        assertEquals(
            parseIntArray("[2,2,0,2,1,2]").toList(),
            assignTasks(parseIntArray("[3,3,2]"), parseIntArray("[1,2,3,2,1,2]")).toList()
        )
        assertEquals(
            parseIntArray("[1,4,1,4,1,3,2]").toList(),
            assignTasks(parseIntArray("[5,1,4,3,2]"), parseIntArray("[2,1,2,4,5,2,1]")).toList()
        )
    }
}
