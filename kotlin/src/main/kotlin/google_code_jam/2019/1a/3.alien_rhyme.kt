package google_code_jam.`2019`.`1a`
fun main(){
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull{if(it != "") it.toInt() else null}
    fun input_longs() = split_input().mapNotNull{if(it != "") it.toLong() else null}


    class Node{
        var cnt = 0
        val ch = hashMapOf<Char, Node>()
    }
    val T = readLine()!!.toInt()
    (1..T).forEach {
        val (N, ) = input_ints()
        val root = Node()
        for(_i in 0 until N){
            var u = root
            u.cnt ++
            for(ch in split_input()[0].reversed()){
                if(ch !in u.ch)
                    u.ch[ch] = Node()
                u = u.ch[ch]!!
                u.cnt ++
            }
        }
        fun dfs(u: Node): Int{
            var done = 0
            for((_c, v) in u.ch){
                done += dfs(v)
            }
            if(u != root && u.cnt - done >= 2){
                done += 2
            }
            return done
        }
        val done = dfs(root)

        println("Case #${it}: $done")
    }
}
