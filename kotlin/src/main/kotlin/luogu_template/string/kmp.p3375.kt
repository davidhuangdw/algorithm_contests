package luogu_template.string

open class KMP(val pattern: String) {
    val n = pattern.length
    val preLen = IntArray(pattern.length)
    init {
        for ((i, ch) in pattern.withIndex().drop(1)){
            var j = preLen[i-1]
            while(j > 0 && pattern[j] != ch){
                j = preLen[j-1]
            }
            preLen[i] = if(pattern[j]==ch) j+1 else 0
        }
    }

    fun findAll(s: String): List<Int>{
        val ind = mutableListOf<Int>()
        var j = 0
        for((i, ch) in s.withIndex()){
            while(j > 0 && pattern[j] != ch){
                j = preLen[j-1]
            }
            j = if(pattern[j] == ch) j+1 else 0
            if(j == n){
                ind.add(i+1-n)
                j = preLen[j-1]
            }
        }
        return ind
    }
}

fun main() {
    val s1 = readLine()!!
    val kmp = KMP(readLine()!!)
    for(i in kmp.findAll(s1)){
        println(i+1)
    }
    println(kmp.preLen.joinToString(" "))
}
