package leetcode

import org.junit.jupiter.api.Assertions.assertArrayEquals
import org.junit.jupiter.api.Test

fun parseArray(s: String): List<String> {
    val input = s.trim()
    val subs = mutableListOf<String>()
    var extra = 0
    var l = 0
    for ((i, ch) in input.withIndex()) {
        when (ch) {
            '[' -> {
                extra += 1
                if (extra == 1) l = i + 1
            }
            ']' -> {
                extra -= 1
                if (extra == 0) {
                    subs.add(input.substring(l, i).trim())
                    l = i + 1
                }
            }
            ',' -> if (extra == 1) {
                subs.add(input.substring(l, i).trim())
                l = i + 1
            }
        }
    }
    return subs
}

inline fun <reified T> parseTypedArray(s: String, convert: (String) -> T) =
    parseArray(s).map { convert(it) }.toTypedArray()


fun parseIntArray(s: String) = parseTypedArray(s) { it.toInt() }.toIntArray()
fun parse2dIntArray(s: String) = parseTypedArray(s) { parseIntArray(it) }

fun parseLongArray(s: String) = parseTypedArray(s) { it.toLong() }.toLongArray()
fun parse2dLongArray(s: String) = parseTypedArray(s) { parseLongArray(it) }

fun parseDoubleArray(s: String) = parseTypedArray(s) { it.toDouble() }.toDoubleArray()
fun parse2dDoubleArray(s: String) = parseTypedArray(s) { parseDoubleArray(it) }

class TestLeetUtils {
    @Test
    fun test1() {
        assertArrayEquals(intArrayOf(1, 2, 3, 4), parseIntArray("  [1, 2, 3, 4] "))

        val input = " [ [0, 1, 2], [13, 14, 15 ], [16, 17, 18] ] "
        var expect = arrayOf(
            intArrayOf(0, 1, 2),
            intArrayOf(13, 14, 15),
            intArrayOf(16, 17, 18),
        )
        for ((a, b) in expect.zip(parse2dIntArray(input)))
            assertArrayEquals(a, b)
    }
}
