package atcoder.a1

fun main() {
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull { if (it != "") it.toInt() else null }
    val (H, W) = input_ints()
    val mp = (1..H).map { readLine()!!.toMutableList() }
    val colors = "12345"
    for (i in 0 until H)
        for (j in 0 until W)
            if(mp[i][j] == '.'){
                mp[i][j] = colors.find{c ->
                    !listOf(i+1 to j, i-1 to j, i to j-1, i to j+1).any{(a,b) ->
                        0<=a && a<H && 0<=b && b<W && mp[a][b]==c
                    }
                }!!
            }
    for(row in mp){
        println(row.joinToString(""))
    }
}