require("dotenv").config();
var MongoClient = require("mongodb").MongoClient;
const ObjectId = require("mongodb").ObjectId
const ids = require("./constant")
const COLLECTION_NAME = process.env.COLLECTION_NAME

const createDbObject = () => {
    let client = null;
    let db = null;

    const connect = async () => {
        if (!db) {
            const url = process.env.MONGO_URL;
            console.info("Connecting to database...");
            client = await MongoClient.connect(url);
            db = client.db("application-prod-db");
            console.info("Connected to database...");
        }
        return db;
    };

    return {
        forLoopExecute: async function () {
            const database = await connect();
            const data = []
            for(const id of ids) {
                const query = {_id: new ObjectId(id)};
                const queryResult = await database.collection(COLLECTION_NAME)
                .find(query)
                .toArray()
                data.push(queryResult)
            }        
        },
        promiseExecute: async function () {
            const database = await connect();
            return await Promise.all(ids.map((id) => {
                const query = {_id: new ObjectId(id)};
                return database.collection(COLLECTION_NAME)
                .find(query)
                .toArray()
            }))   
        },
        bulkExecute: async function () {
            const database = await connect();
            const query = {_id: {"$in": ids.map((id) => new ObjectId(id))}};
            return await database.collection(COLLECTION_NAME)
                .find(query)
                .toArray();
        },
        close: async function () {
            if (client) {
                await client.close();
                client = null;
                db = null;
                console.info("Database connection closed.");
            }
        }
    };
};

async function init() {
    try {
        const functionName = process.argv[2]; 

        const db = createDbObject();
        
        if (!db?.[functionName]) {
            throw `Error: Function '${functionName}' not found!`;
        }

        console.time(`Execution Time (${functionName}): `)
        const beforeMemory = process.memoryUsage();
        await db[functionName]();
        const afterMemory = process.memoryUsage();
        console.timeEnd(`Execution Time (${functionName}): `)
        console.info(`Memory spike (${functionName}):", ${afterMemory.heapUsed - beforeMemory.heapUsed} bytes`);
        await db.close();
    } catch(e) {
        console.error("Error: ", e);
    }
}

init()
